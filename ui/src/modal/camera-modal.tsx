import React from "react";

import { Device } from "../interfaces/device.interface";

import "./camera-modal.scss";

type Props = {
  device: Device;
};

export const CameraModal: React.FC<Props> = (props) => {
  return (
    <div className="modal-camera">
      <img src={`http://${props.device.ip}:81/stream`} />
    </div>
  );
};
