import React from "react";

import "./thermostat.scss";
import { LightbulbFill, LightbulbOff } from "react-bootstrap-icons";

import { AlertState, Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
  onOpenModal: (device: Device) => void;
};

export const Camera: React.FC<Props> = (props: Props) => {
  return (
    <section
      className={`tile tile-type-plug tile-state-${alert}`}
      onClick={() => props.onOpenModal(props.device)}
    >
      <h2>{props.device.name}</h2>
      <div className="tile-body">
        <div>
          {props.device.state.alert === AlertState.ON ? (
            <LightbulbFill />
          ) : (
            <LightbulbOff />
          )}
        </div>
      </div>
    </section>
  );
};
