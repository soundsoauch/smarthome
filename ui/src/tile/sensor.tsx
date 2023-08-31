import React from "react";
import { UnlockFill, Lock } from "react-bootstrap-icons";

import { AlertState, Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
};

export const Sensor: React.FC<Props> = (props: Props) => {
  return (
    <div className="tile-body">
      {props.device.state.alert === AlertState.OPEN ? <UnlockFill /> : <Lock />}
    </div>
  );
};
