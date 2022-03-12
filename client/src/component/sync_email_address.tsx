import React from 'react';
import { Something } from '../types/something';

export default function SyncEmailAddress({ somethings,syncEmailAddressVal }): JSX.Element {
  return somethings.map((something: Something) => {
      syncEmailAddressVal.current.value = something.appEmailAddress
      if (something.type != 'recv') {
          return ('');
      }
  });
}