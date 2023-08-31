import React from "react";
import { useState } from "react";
import Card from "react-bootstrap/Card";
import Stack from "react-bootstrap/Stack";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Container from "react-bootstrap/Container";

import {
  UnlockFill,
  Lock,
  ThermometerHalf,
  Water,
  LightbulbFill,
  LightbulbOff,
} from "react-bootstrap-icons";

import { Device, DeviceType, AlertState } from "../interfaces/device.interface";

import "./tile.scss";

type Props = {
  device: Device;
};

export const Tile: React.FC<Props> = (props: Props) => {
  const template: React.FC = (device: Device) => {
    switch (device.type) {
      case DeviceType.THERMOMETER:
        return (
          <div className="tile-body">
            <div>
              {device.state.temperature / 10}
              <div className="title-unit">
                <em>°C</em>
                <ThermometerHalf />
              </div>
            </div>
            <div>
              {device.state.humidity}
              <div className="title-unit">
                <em>%</em>
                <Water />
              </div>
            </div>
          </div>
        );
      case DeviceType.SENSOR:
        return (
          <div className="tile-body">
            {device.state.alert === AlertState.OPEN ? <UnlockFill /> : <Lock />}
          </div>
        );
      case DeviceType.THERMOSTAT:
        return (
          <div className="tile-body">
            <div>
              {device.state.temperature / 10}
              <div className="title-unit">
                <em>°C</em>
                <ThermometerHalf />
              </div>
            </div>
          </div>
        );
      case DeviceType.PLUG:
        return (
          <div className="tile-body">
            <div>
              {device.state.alert === AlertState.ON ? (
                <LightbulbFill />
              ) : (
                <LightbulbOff />
              )}
            </div>
          </div>
        );
    }
    return <></>;
  };

  const classNames: string[] = ["tile", "tile-type-" + props.device.type];

  return (
    <section className={classNames.join(" ")}>
      <h2>{props.device.name}</h2>
      {template(props.device)}
    </section>
  );
};
