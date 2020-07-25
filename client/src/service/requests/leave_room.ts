import Service from '../service';
import { C2S } from '../../models';

export default function leaveRoom(service: Service) {
  service.send({
    cmd: C2S.CLIENT.LEAVE_ROOM,
    data: {},
  });
}
