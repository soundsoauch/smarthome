import React, { useState } from "react";

import "./modal.scss";
import { Device, StateDto, DeviceType } from "../interfaces/device.interface";
import { ThermostatModal } from "./thermostat-modal";
import { CameraModal } from "./camera-modal";

type Props = {
  device: Device;
  onClose: () => void;
  onUpdate: (id: string, deviceState: StateDto) => void;
};

export const Modal: React.FC<Props> = (props) => {
  const type = props.device.type;
  const getTemplate = () => {
    switch (type) {
      case DeviceType.THERMOSTAT:
        return (
          <ThermostatModal device={props.device} onUpdate={props.onUpdate} />
        );
      case DeviceType.CAMERA:
        return <CameraModal device={props.device} />;
    }
  };

  return (
    <>
      <div className="modal-stage" onClick={props.onClose}></div>
      <section className="modal">
        <h2>{props.device.name}</h2>
        {getTemplate()}
      </section>
    </>
  );
};
