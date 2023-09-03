import React from "react";

import "./tile.scss";

import { Plug } from "./plug";
import { Device, DeviceType } from "../interfaces/device.interface";
import { Thermostat } from "./thermostat";
import { Thermometer } from "./thermometer";
import { Sensor } from "./sensor";

type Props = {
  device: Device;
  onOpenModal: (device: Device) => void;
};

export const Tile: React.FC<Props> = (props: Props) => {
  const template: React.FC = (device: Device) => {
    switch (device.type) {
      case DeviceType.THERMOMETER:
        return <Thermometer device={device} />;
      case DeviceType.SENSOR:
        return <Sensor device={device} />;
      case DeviceType.THERMOSTAT:
        return <Thermostat device={device} onOpenModal={props.onOpenModal} />;
      case DeviceType.PLUG:
        return <Plug device={device} />;
    }
    return <></>;
  };

  return template(props.device);
};
