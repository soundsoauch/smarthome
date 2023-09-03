import React, { useState } from "react";

import "./modal.scss";
import { Device, StateDto } from "../interfaces/device.interface";

type Props = {
  device: Device;
  onClose: () => void;
  onUpdate: (id: string, deviceState: StateDto) => void;
};

interface FormData {
  tSoll?: number;
}

export const Modal: React.FC<Props> = (props) => {
  const id = props.device.id;
  const type = props.device.type;
  const [value, setValue] = useState<FormData>({
    tSoll: props.device.state?.tSoll ?? 75,
  });

  const targetChange: (e: React.ChangeEvent<HTMLInputElement>) => void = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setValue({ tSoll: Number(e.target.value) * 10 });
  };
  const submitForm: (e: React.FormEvent) => void = (e: React.FormEvent) => {
    e.preventDefault();

    fetch("/devices?type=" + type + "&id=" + id, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        tSoll: value.tSoll,
      }),
    })
      .then((res: Response) => {
        if (res.status === 200) {
          props.onUpdate(props.device.id, value);
        }
      })
      .catch((err: Error) => {
        console.error(err.message);
      });
  };

  return (
    <>
      <div className="modal-stage" onClick={props.onClose}></div>
      <section className="modal">
        <h2>{props.device.name}</h2>
        <form onSubmit={(e: React.FormEvent) => submitForm(e)}>
          <div>
            <label htmlFor="tSoll">
              Target Temperature
              <br />({value?.tSoll > 75 ? value.tSoll / 10 : "OFF"})
            </label>
            <input
              id="tSoll"
              name="tSoll"
              type="range"
              min="7.5"
              max="28"
              step="0.5"
              value={value?.tSoll / 10}
              onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
                targetChange(e)
              }
            />
          </div>
          <input type="submit" value="Submit" />
        </form>
      </section>
    </>
  );
};
