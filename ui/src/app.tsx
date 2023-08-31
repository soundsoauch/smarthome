import React from "react";
import { useState, useEffect } from "react";

import Container from "react-bootstrap/Container";
import Navbar from "react-bootstrap/Navbar";

import "./app.scss";
import { Tile } from "./tile/tile";
import { Device, DeviceType } from "./interfaces/device.interface";
import { GetDevices } from "./hooks/data";

export default function App() {
  const devices: Device[] = GetDevices();

  return (
    <>
      <Navbar sticky="top" className="bg-body-tertiary">
        <Container>
          <Navbar.Brand href="#home">Smarthome</Navbar.Brand>
        </Container>
      </Navbar>
      <ul className="devices-container">
        {devices.map((device: Device) => (
          <li key={device.id}>
            <Tile device={device} />
          </li>
        ))}
      </ul>
    </>
  );
}
