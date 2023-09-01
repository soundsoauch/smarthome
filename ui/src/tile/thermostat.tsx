import React from "react";
import { ThermometerHalf } from "react-bootstrap-icons";

import { Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
};

export const Thermostat: React.FC<Props> = (props: Props) => {
  return (
    <section className="tile tile-type-thermostat">
      <h2>{props.device.name}</h2>
      <div className="tile-body">
        <div>
          {props.device.state.temperature / 10}
          <div className="title-unit">
            <em>°C</em>
            <ThermometerHalf />
          </div>
        </div>
      </div>
    </section>
  );
};
