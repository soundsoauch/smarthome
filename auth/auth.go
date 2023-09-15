package auth

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"smarthome/interfaces"
	"unicode/utf16"

	"github.com/gin-gonic/gin"
)

type Session struct {
	host       string
	password   string
	username   string
	currentSID *string
}

type SessionInfo struct {
	Challenge string `xml:"Challenge"`
	SID       string `xml:"SID"`
	BlockTime string `xml:"BlockTime"`
}

func New(host string, username string, password string) *Session {
	return &Session{host: host, username: username, password: password}
}

func (s *Session) GetSID() (string, error) {
	if s.currentSID == nil {
		return s.Login()
	}
	if *s.currentSID == "" {
		return s.Login()
	}
	if isValid, _ := s.isSIDValid(); !isValid {
		return s.Login()
	}

	return *s.currentSID, nil
}

func (s *Session) isSIDValid() (bool, error) {
	resp, err := http.Get(s.host + "/login_sid.lua?sid=" + *s.currentSID)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var sessionResp SessionInfo
	err = xml.Unmarshal(body, &sessionResp)
	if err != nil {
		return false, err
	}

	if sessionResp.SID != "0000000000000000" {
		return true, nil
	}

	return false, nil
}

func (s *Session) GetChallange() (string, error) {
	resp, err := http.Get(s.host + "/login_sid.lua")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var prelogin SessionInfo
	err = xml.Unmarshal(body, &prelogin)
	if err != nil {
		return "", err
	}

	return prelogin.Challenge, nil
}

func (s *Session) Login() (string, error) {
	challange, err := s.GetChallange()
	if err != nil {
		return "", err
	}

	resp, err := http.Get(fmt.Sprintf(
		"%s/login_sid.lua?response=%s-%s&username=%s",
		s.host, challange, preparePassword(challange, s.password), s.username,
	))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	var login SessionInfo
	err = xml.Unmarshal(body, &login)
	if err != nil {
		return "", err
	}

	if login.SID != "0000000000000000" {
		s.currentSID = &login.SID
		return login.SID, nil
	} else {
		return "", errors.New("failed to login, try again in " + login.BlockTime + " second(s)")
	}
}

func preparePassword(challenge, password string) string {
	converted := utf16.Encode([]rune(challenge + "-" + password))
	b := make([]byte, 2*len(converted))
	for i, v := range converted {
		binary.LittleEndian.PutUint16(b[i*2:], v)
	}

	hasher := md5.New()
	hasher.Write(b)
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func Authenticator(config *interfaces.Config) gin.HandlerFunc {
	var session *Session = New(config.Fritzbox.Host, config.Fritzbox.User, config.Fritzbox.Password)

	return func(c *gin.Context) {
		var sid, err = session.GetSID()

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.Set("sid", sid)
		c.Next()
	}
}