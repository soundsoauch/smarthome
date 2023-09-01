import React, { useState } from "react";

import "./plug.scss";

import { LightbulbFill, LightbulbOff } from "react-bootstrap-icons";

import { AlertState, Device } from "../interfaces/device.interface";

type Props = {
  device: Device;
};

export const Plug: React.FC<Props> = (props: Props) => {
  const [alert, setAlert] = useState<AlertState>(props.device.state.alert);
  const id = props.device.id;
  const type = props.device.type;

  const toggleState: () => void = () => {
    const nextState: AlertState =
      alert === AlertState.ON ? AlertState.OFF : AlertState.ON;

    fetch("/devices?type=" + type + "&id=" + id, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        alert: nextState,
      }),
    })
      .then((res: Response) => {
        if (res.status === 200) {
          setAlert(nextState);
        }
      })
      .catch((err: Error) => {
        console.error(err.message);
      });
  };

  return (
    <section
      className={`tile tile-type-plug tile-state-${alert}`}
      onClick={() => toggleState()}
    >
      <h2>{props.device.name}</h2>
      <div className="tile-body">
        <div>
          {alert === AlertState.ON ? <LightbulbFill /> : <LightbulbOff />}
        </div>
      </div>
    </section>
  );
};
