import React from 'react';
import { Something } from '../types/something';

export default function SyncUsername({ somethings,syncUsernameVal }): JSX.Element {
  return somethings.map((something: Something) => {
      syncUsernameVal.current.value = something.appUsername
      if (something.type != 'recv') {
          return ('');
      }
  });
}