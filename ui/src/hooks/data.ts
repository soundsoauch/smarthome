import { useState, useEffect } from "react";

import { Device, DeviceType, StateDto } from "../interfaces/device.interface";

export const GetDevices: () => Device[] = () => {
  const [devices, setDevices] = useState<Device[]>([]);

  useEffect(() => {
    fetch("/devices")
      .then((res: Response) => res.json())
      .then((devices: Device[]) =>
        devices.filter((device: Device) => device.type !== DeviceType.BUTTON)
      )
      .then((devices: Device[]) => setDevices(devices))
      .catch((err: Error) => {
        console.error(err.message);
        return [];
      });
  }, []);

  return devices;
};

export const SetDevice: (
  id: string,
  type: DeviceType,
  state: StateDto
) => boolean = (id: string, type: DeviceType, state: StateDto) => {
  const [response, setResponse] = useState<boolean>(null);

  useEffect(() => {
    fetch("/devices?type=" + type + "&id=" + id, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(state),
    })
      .then((res: Response) => {
        setResponse(res.status === 200);
      })
      .catch((err: Error) => {
        console.error(err.message);
        return false;
      });
  }, []);

  return response;
};
