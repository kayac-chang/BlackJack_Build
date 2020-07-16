import { S2C } from '../../models';
import Service from '../service';

function onError(service: Service, data: any) {
  throw new Error(JSON.stringify(data));
}

export default {
  [S2C.SYSTEM.ERROR]: onError,
};
