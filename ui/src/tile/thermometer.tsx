import React from "react";
import { ThermometerHalf, Water } from "react-bootstrap-icons";

import { Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
};

export const Thermometer: React.FC<Props> = (props: Props) => {
  return (
    <div className="tile-body">
      <div>
        {props.device.state.temperature / 10}
        <div className="title-unit">
          <em>°C</em>
          <ThermometerHalf />
        </div>
      </div>
      <div>
        {props.device.state.humidity}
        <div className="title-unit">
          <em>%</em>
          <Water />
        </div>
      </div>
    </div>
  );
};
