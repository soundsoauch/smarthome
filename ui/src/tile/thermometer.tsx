import React from "react";
import { ThermometerHalf, Water } from "react-bootstrap-icons";

import { Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
};

export const Thermometer: React.FC<Props> = (props: Props) => {
  return (
    <section className="tile tile-type-thermometer">
      <h2>{props.device.name}</h2>
      <div className="tile-body">
        <div>
          {props.device.state.temperature / 10}
          <div className="title-unit">
            <em>°C</em>T
          </div>
        </div>
        <div>
          {props.device.state.humidity}
          <div className="title-unit">
            <em>%</em>
            RH
          </div>
        </div>
      </div>
    </section>
  );
};
