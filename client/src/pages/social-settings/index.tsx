import React, { useContext, useEffect, useRef, useState } from 'react';
import { Range, getTrackBackground } from 'react-range';
import { WebSocketContext } from '../../modules/websocket_provider';
import router from 'next/router';
import SyncUsername from "../../component/sync_username";
import SyncEmailAddress from "../../component/sync_email_address";
import { AuthContext } from '../../modules/auth_provider';
import { Something } from '../../types/something';
import { useGetDevices } from '../../hooks/use_get_devices';
import { getDevicesInUser } from '../../service/get_devices_in_user'
import Loading from '../../component/loading';

export default function SocialSettings() {
  const [somethings, setSomethings] = useState<Array<Something>>([]);
  const usernameVal = useRef(null);
  const emailAddressVal = useRef(null);
  const alignedCbCheck = useRef(null);
  const [billingPeriodVal, setBillingPeriodVal] = useState({values: [3]});
  const syncBillingPeriodVal = useRef({values: [3]});
  const [salaryVal, setSalaryVal] = useState({values: [2500]});
  const syncSalaryVal = useRef({values: [2500]});
  const [alignedCbVal, setAlignedCbVal] = useState({checked: false});
  const syncAlignedCbVal = useRef({checked: false});
  const { conn, setConn } = useContext(WebSocketContext);
  const { jwtClaims } = useContext(AuthContext);
  const [joinStatus, setJoinStatus] = useState(''); // TODO: Status joined device(s) = (joined, disjoined, rejoin)  ( joined_device, disjoined_device )
  const { devices, setDevices } = useGetDevices(); // TODO: call 'useGetDevices' >> 'getDevicesInUser' from join user with 'username'

  useEffect(() => {
    console.log(" ...app.useEffect: jwtClaims <<<", jwtClaims);

    if (conn === null) {
      router.push('/');
      return;
    }

    conn.onclose = (closeEvent) => {
        setJoinStatus('disjoined');
    };

    conn.onerror = (event) => {
        setJoinStatus('joining error');
    };

    conn.onopen = (event) => {
        const pathname = new URL(conn.url).pathname.split('/');
        setJoinStatus('joined ' + pathname[2]); // TODO: username
    };

    conn.onmessage = (messageEvent) => { // TODO: receive remote Something(s)
      const something: Something = JSON.parse(messageEvent.data);

        if (something.appUsername == 'joined_device') {
            setDevices([...devices, { deviceName: something.deviceName }]); // TODO: [one special] sync joined device(s) by user
            return;
        }
        if (something.appUsername == 'disjoined_device') {
            const deleteDevice = devices.filter((device) => device.deviceName != something.deviceName); // TODO: [one special] sync disjoined device(s) by user
            setDevices([...deleteDevice]);
            return;
        }
      jwtClaims.id == something.id ? (something.type = 'recv') : (something.type = 'self');

      setAlignedCbVal({checked: something.appAlignedCb});
      setBillingPeriodVal({values: [something.appBillingPeriod]});
      setSalaryVal({values: [something.appSalary]});
      syncAlignedCbVal.current.checked = something.appAlignedCb;
      syncBillingPeriodVal.current.values = [something.appBillingPeriod];
      syncSalaryVal.current.values = [something.appSalary];

      setSomethings([...somethings, something]);
      console.log(' ...app.useEffect: conn.onmessage (somethings) <<<', somethings)
    }
  }, [conn, devices, somethings, usernameVal, emailAddressVal]);

  const sendSomething = () => {
    let data = {
      appUsername: usernameVal.current.value,
      appEmailAddress: emailAddressVal.current.value,
      appAlignedCb: alignedCbCheck.current.checked,
      appBillingPeriod: syncBillingPeriodVal.current.values[0],
      appSalary: syncSalaryVal.current.values[0],
    }

    console.log(' ...app.send: data (somethings) >>>', data);
    conn.send( JSON.stringify(data) ); // TODO: send locale Something(s)
  };

  const rejoin = () => {
    if (conn == null) {
      return router.push('/');
    }
    const ws = new WebSocket(conn.url);
    if (ws.OPEN) {
      setConn(ws);
      // setDevices([]);
      const pathname = conn.url.split('/')
      getDevicesInUser(pathname[4]) // TODO: username
        .then((res) => {
            console.log(' ...use_get_devices.getDevicesInUser <<<', res.data); // TODO: call 'getDevicesInUser' from join user with 'username'
            setDevices(res.data.data);
        });
    }
  };

  if (devices === [] || conn === null) <Loading />;

  console.log(JSON.stringify(" ...app: devices <<<", devices))

  return (
      <div className="flex flex-col md:flex-row w-full">
        <div className="flex items-center justify-center w-full md:w-9/14">
          <div className="p-4 md:mx-24 mb-14">
            <div className="flex items-center justify-center"
                 style={{
                     fontSize: "34px",
                     fontWeight: "bold"
                 }}>
                Social server | <small>settings sync demo</small>
            </div>

            <div className="mt-6 grid grid-cols-6">
                <div />
                <div className="text text-dark-secondary text-right  md:w-6/6 flex flex-col border-b-2 border-dark-secondary">
                    <legend>User form (block 1 to sync)</legend>
                </div>
                <div className="md:w-6/6 flex flex-col border-b-2 border-dark-secondary" />
                <div className="md:w-6/6 flex flex-col border-b-2 border-dark-secondary" />
                <div className="md:w-6/6 flex flex-col border-b-2 border-dark-secondary" />
            </div>
            <div className="mt-2 grid grid-cols-4">
                <div />
                <div className="text-right pt-1">
                    <legend>Username</legend>
                </div>
                <div className="pl-3" style={{borderRadius: '5px'}}>
                    <input className="p-2 focus:outline-none border border-dark-primary focus:border-blue"
                           ref={usernameVal}
                           style={{
                               borderRadius: '5px',
                               width: '250px',
                               height: '36px'
                           }}
                           type="text"
                           placeholder="Username"
                           onChange={sendSomething}/>
                    <SyncUsername somethings={somethings} syncUsernameVal={usernameVal} />
                </div>
                <div />
            </div>
            <div className="mt-2 grid grid-cols-4">
                <div />
                <div className="text-right pt-1">
                  <legend>Email Address</legend>
                </div>
                <div className="pl-3" style={{borderRadius: '5px'}}>
                    <input className="p-2 focus:outline-none border border-dark-primary focus:border-blue"
                           ref={emailAddressVal}
                           style={{
                               borderRadius: '5px',
                               width: '250px',
                               height: '36px'
                           }}
                           type="text"
                           placeholder="Email Address"
                           onChange={sendSomething}/>
                    <SyncEmailAddress somethings={somethings} syncEmailAddressVal={emailAddressVal} />
                </div>
                <div />
            </div>

            <div className="mt-2 grid grid-cols-4">
                <div />
                <div />
                <div className="pl-3" style={{borderRadius: '5px'}}>
                    <input type="checkbox"
                           ref={alignedCbCheck}
                           onChange={sendSomething}
                           checked={syncAlignedCbVal.current.checked}
                            /> I&#x27;ve read the terms and conditions
                </div>
                <div />
            </div>
            <div className="mt-2 grid grid-cols-6">
                <div />
                <div className="text text-dark-secondary text-right  md:w-6/6 flex flex-col border-b-2 border-dark-secondary">
                    <legend>Billing form (block 2 to sync)</legend>
                </div>
                <div className="md:w-6/6 flex flex-col border-b-2 border-dark-secondary" />
                <div className="md:w-6/6 flex flex-col border-b-2 border-dark-secondary" />
                <div className="md:w-6/6 flex flex-col border-b-2 border-dark-secondary" />
            </div>
            <div className="mt-2 grid grid-cols-6">
                <div />
                <div className="text-right">
                    <label>Billing period {billingPeriodVal.values} month</label>
                </div>
                <div className="pt-2.5 pl-3">
                    <Range
                        min={1}
                        max={12}
                        step={1}
                        values={billingPeriodVal.values}

                        onChange={
                            (values) => {
                                setBillingPeriodVal({values});
                                syncBillingPeriodVal.current.values = values;
                                sendSomething();
                            }
                        }

                        renderTrack={({props, children}) => (
                            <div className="flex w-full mr-4 border border-dark-secondary"
                                {...props}
                                onMouseDown={props.onMouseDown}
                                style={{
                                    ...props.style,
                                    width: "200px",
                                    height: "6px",
                                    borderRadius: "6px",
                                    background: getTrackBackground({
                                        values: billingPeriodVal.values,
                                        colors: ["#e95420", "#e8e8e8"],
                                        min: 1,
                                        max: 12
                                    }),
                                    cursor: "default",
                                    display: "flex",
                                }}>
                                {children}
                            </div>
                        )}

                        renderThumb={({props, isDragged}) => (
                            <div
                                {...props}
                                style={{
                                    ...props.style,
                                    width: "20px",
                                    height: "20px",
                                    borderRadius: "20px",
                                    borderWidth: '2px',
                                    borderColor: '#FFF',
                                    backgroundColor: isDragged ? "#e95420" : "#616161",
                                    boxShadow: "0px 2px 2px #AAA",
                                    cursor: "default",
                                    justifyContent: "center",
                                    alignItems: "center",
                                    display: "flex",
                                }}>
                            </div>
                        )}
                    />
                </div>

                <div className="text-right">
                    <label>Salary ${salaryVal.values}</label>
                </div>
                <div className="pt-2.5 pl-5">
                    <Range
                        min={0}
                        max={10000}
                        step={1}
                        values={salaryVal.values}

                        onChange={
                            (values) => {
                                setSalaryVal({values});
                                syncSalaryVal.current.values = values;
                                sendSomething();
                            }
                        }

                        renderTrack={({props, children}) => (
                            <div className="flex w-full mr-4 border border-dark-secondary"
                                 {...props}
                                 onMouseDown={props.onMouseDown}
                                 style={{
                                     ...props.style,
                                     width: "200px",
                                     height: "6px",
                                     borderRadius: "6px",
                                     background: getTrackBackground({
                                         values: salaryVal.values,
                                         colors: ["#e95420", "#e8e8e8"],
                                         min: 0,
                                         max: 10000
                                     }),
                                     cursor: "default",
                                     display: "flex",
                                 }}>
                                {children}
                            </div>
                        )}

                        renderThumb={({props, isDragged}) => (
                            <div
                                {...props}
                                style={{
                                    ...props.style,
                                    width: "20px",
                                    height: "20px",
                                    borderRadius: "20px",
                                    borderWidth: '2px',
                                    borderColor: '#FFF',
                                    backgroundColor: isDragged ? "#e95420" : "#616161",
                                    boxShadow: "0px 2px 2px #AAA",
                                    cursor: "default",
                                    justifyContent: "center",
                                    alignItems: "center",
                                    display: "flex",
                                }}>
                            </div>
                        )}
                    />
                </div>
                <div />
            </div>
          </div>
        </div>

        <div className="md:w-3/12 md:visible invisible flex flex-col border-l-2 border-dark-secondary p-4">
          <div className="fixed">
            <OnCloseJoin rejoin={rejoin} joinStatus={joinStatus} />
            <div className="mb-4 text-lg font-bold">sync device(s)</div>
            {devices.map((device, index) => (
              <div className="ml-4 flex flex-row items-center h-full min-w-full" key={index}>
                <div className="h-2 w-2 mr-2 bg-green  items-center rounded-full"></div>
                <div>{device.deviceName}</div>
              </div>
            ))}
          </div>
        </div>
      </div>
  );
}

export function OnCloseJoin({ rejoin, joinStatus }) {
  const joinedStyleErr = 'p-2 px-4 flex flex-row justify-end w-full  bg-red bg-opacity-10 text-red rounded-md';
  const joinedStyleOn =  'p-2 px-4 flex flex-row justify-end w-full  bg-green bg-opacity-10 text-green rounded-md';

  return (
    <div className="mb-4 inline-block">
      <div className={ joinStatus.includes('disjoined') ? joinedStyleErr: joinedStyleOn }>
        <div>{joinStatus}</div>
        {joinStatus.includes('disjoined') && (
          <div>
            <button className="px-5 ml-5  btn btn-outline-dark btn-sm  border border-red rounded-md"
                    type="button"
                    onClick={rejoin}>
                rejoin
            </button>
          </div>
        )}
      </div>
    </div>
  );
}
