import React from "react";
import { UnlockFill, Lock } from "react-bootstrap-icons";

import "./sensor.scss";

import { AlertState, Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
};

export const Sensor: React.FC<Props> = (props: Props) => {
  return (
    <section
      className={`tile tile-type-sensor tile-state-${props.device.state.alert}`}
    >
      <h2>{props.device.name}</h2>
      <div className="tile-body">
        {props.device.state.alert === AlertState.OPEN ? (
          <UnlockFill />
        ) : (
          <Lock />
        )}
      </div>
    </section>
  );
};
