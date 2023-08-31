import React from "react";
import { LightbulbFill, LightbulbOff } from "react-bootstrap-icons";

import { AlertState, Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
};

export const Plug: React.FC<Props> = (props: Props) => {
  return (
    <div className="tile-body">
      <div>
        {props.device.state.alert === AlertState.ON ? (
          <LightbulbFill />
        ) : (
          <LightbulbOff />
        )}
      </div>
    </div>
  );
};
