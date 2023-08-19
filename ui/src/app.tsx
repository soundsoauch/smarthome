import React from "react";
import { useState, useEffect } from "react";

import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";
import Card from "react-bootstrap/Card";

import "./app.scss";

export function App() {
  const [devices, setDevices] = useState([]);
  const [switches, setSwitches] = useState([]);

  useEffect(() => {
    fetch("/devices")
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setDevices(data);
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);
  useEffect(() => {
    fetch("/switches")
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        setSwitches(data);
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  return (
    <>
      <Navbar className="bg-body-tertiary">
        <Container>
          <Navbar.Brand href="#home">Smarthome</Navbar.Brand>
        </Container>
      </Navbar>
      <Card>
        <Card.Body>
          <pre>{JSON.stringify(devices, null, 2)}</pre>
        </Card.Body>
      </Card>
      <Card>
        <Card.Body>
          <pre>{JSON.stringify(switches, null, 2)}</pre>
        </Card.Body>
      </Card>
    </>
  );
}
