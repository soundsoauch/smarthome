import React, { useState, useEffect } from "react";

import "./app.scss";
import { Tile } from "./tile/tile";
import { Device, DeviceType, StateDto } from "./interfaces/device.interface";
import { Modal } from "./modal/modal";

export default function App() {
  const [devices, setDevices] = useState<Device[]>([]);
  const [modal, setModal] = useState<Device>();

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

  const onOpenModal = (device: Device) => {
    setModal(device);
  };

  const onUpdate: (id: string, deviceState: StateDto) => void = (
    id: string,
    deviceState: StateDto
  ) => {
    setDevices(
      devices.map((device: Device) => {
        if (device.id === id) {
          const getTSoll = (tSoll: number) => (tSoll < 80 ? null : tSoll);

          return {
            ...device,
            state: {
              ...device.state,
              ...(deviceState.tSoll
                ? { tSoll: getTSoll(deviceState.tSoll) }
                : {}),
              ...(deviceState.alert ? { alert: deviceState.alert } : {}),
            },
          };
        }
        return device;
      })
    );
    setModal(null);
  };

  return (
    <>
      <ul className="devices-container">
        {devices.map((device: Device) => (
          <li key={device.id}>
            <Tile
              device={device}
              onOpenModal={(device) => onOpenModal(device)}
            />
          </li>
        ))}
      </ul>
      {modal && (
        <Modal
          device={modal}
          onUpdate={(id: string, deviceState: StateDto) =>
            onUpdate(id, deviceState)
          }
          onClose={() => setModal(null)}
        />
      )}
    </>
  );
}
