import React from "react";
import { ThermometerHalf } from "react-bootstrap-icons";

import "./thermostat.scss";

import { Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
  onOpenModal: (device: Device) => void;
};

export const Thermostat: React.FC<Props> = (props: Props) => {
  return (
    <section
      className="tile tile-type-thermostat"
      onClick={() => props.onOpenModal(props.device)}
    >
      <h2>{props.device.name}</h2>
      <div className="tile-body">
        <div>
          {props.device.state.temperature / 10}
          <div className="title-unit">
            <em>°C</em>
            cur
          </div>
        </div>
        {props.device.state.tSoll && (
          <div>
            {props.device.state.tSoll / 10}
            <div className="title-unit">
              <em>°C</em>
              tar
            </div>
          </div>
        )}
      </div>
    </section>
  );
};
