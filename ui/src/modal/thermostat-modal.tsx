import React, { useState } from "react";

import { Device, StateDto, State } from "../interfaces/device.interface";

type Props = {
  device: Device;
  onUpdate: (id: string, deviceState: StateDto) => void;
};

interface FormData {
  tSoll?: number;
}

export const ThermostatModal: React.FC<Props> = (props) => {
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

    fetch("/devices?type=" + props.device.type + "&id=" + props.device.id, {
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

  const isManual: () => boolean = () => {
    const state: State = props.device.state;
    return (
      value.tSoll !== state.absenk &&
      value.tSoll !== state.komfort &&
      value.tSoll !== 75
    );
  };

  const onSetMode: (radioValue: number) => void = (radioValue: number) => {
    setValue({ tSoll: radioValue });
  };

  return (
    <>
      <form onSubmit={(e: React.FormEvent) => submitForm(e)}>
        <div>
          <label htmlFor="tSoll">
            Target Temperature ({value?.tSoll > 75 ? value.tSoll / 10 : "OFF"})
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
        <span>Mode</span>
        <div className="radio-group">
          <div>
            <input
              type="radio"
              name="mode"
              id="off"
              value="7.5"
              checked={value.tSoll === 75}
              onClick={() => onSetMode(75)}
            />
            <label htmlFor="off">OFF</label>
          </div>
          <div>
            <input
              type="radio"
              name="mode"
              id="absenk"
              checked={value.tSoll === props.device.state.absenk}
              value={props.device.state.absenk / 10}
              onClick={() => onSetMode(props.device.state.absenk)}
            />
            <label htmlFor="absenk">Absenk</label>
          </div>
          <div>
            <input
              type="radio"
              name="mode"
              id="komfort"
              checked={value.tSoll === props.device.state.komfort}
              value={props.device.state.komfort / 10}
              onClick={() => onSetMode(props.device.state.komfort)}
            />
            <label htmlFor="komfort">Komfort</label>
          </div>
          <div>
            <input
              type="radio"
              name="mode"
              id="manual"
              value={value?.tSoll / 10}
              disabled
              checked={isManual()}
              className="disabled"
            />
            <label htmlFor="manual" className="disabled">
              Manuel
            </label>
          </div>
        </div>
        <input id="submit-button" type="submit" value="Submit" />
      </form>
    </>
  );
};
