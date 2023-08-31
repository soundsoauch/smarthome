import React from "react";

import "./tile.scss";

import { Plug } from "./plug";
import { Device, DeviceType } from "../interfaces/device.interface";
import { Thermostat } from "./thermostat";
import { Thermometer } from "./thermometer";
import { Sensor } from "./sensor";

type Props = {
  device: Device;
};

export const Tile: React.FC<Props> = (props: Props) => {
  const template: React.FC = (device: Device) => {
    switch (device.type) {
      case DeviceType.THERMOMETER:
        return <Thermometer device={device} />;
      case DeviceType.SENSOR:
        return <Sensor device={device} />;
      case DeviceType.THERMOSTAT:
        return <Thermostat device={device} />;
      case DeviceType.PLUG:
        return <Plug device={device} />;
    }
    return <></>;
  };

  const classNames: string[] = ["tile", "tile-type-" + props.device.type];

  return (
    <section className={classNames.join(" ")}>
      <h2>{props.device.name}</h2>
      {template(props.device)}
    </section>
  );
};
