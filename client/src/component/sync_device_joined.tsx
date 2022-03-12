import React from 'react';
import { Something } from '../types/something';

export default function SyncDeviceJoined({ somethings,syncDeviceJoinedVal }): JSX.Element {
  return somethings.map((something: Something) => {
      syncDeviceJoinedVal.current = something.syncDeviceJoined
      if (something.type != 'recv') {
          return ('');
      }
  });
}