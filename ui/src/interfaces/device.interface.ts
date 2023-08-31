export interface Device {
  name: string;
  id: string;
  isConnected: boolean;
  type?: DeviceType;
  lowBattery?: boolean;
  state?: State;
}

export interface State {
  humidity?: number;
  temperature?: number;
  tSoll?: number;
  alert?: AlertState;
  energy?: Energy;
}

export interface Energy {
  total: number;
  yesterday: number;
  today: number;
  power: number;
}

export interface StateDto {
  tSoll?: number;
  alert?: AlertState;
}

export enum DeviceType {
  THERMOSTAT = "THERMOSTAT",
  SENSOR = "SENSOR",
  PLUG = "PLUG",
  THERMOMETER = "THERMOMETER",
  BUTTON = "BUTTON",
  UNKNOWN = "UNKNOWN",
}

export enum AlertState {
  OPEN = "OPEN",
  CLOSE = "CLOSE",
  ON = "ON",
  OFF = "OFF",
}
