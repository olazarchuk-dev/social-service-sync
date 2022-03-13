export type Something = { // TODO: init dynamic data
  id: number;
  deviceName: string;
  syncDeviceJoined: string;
  deviceBrowser: {};
  lastActiveTime: number;
  type: 'recv' | 'self';
  appUsername: string;
  appEmailAddress: string;
  appAlignedCb: boolean;
  appBillingPeriod: number;
  appSalary: number;
};
