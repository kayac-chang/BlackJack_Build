import EventEmitter from "eventemitter3";
import { login } from "./requests";
import { S2C, C2S, Token } from "../models";
import MUX from "./mux";

interface Frame {
  cmd: S2C.ROOM | S2C.ROUND | S2C.SEAT | S2C.USER | C2S.CLIENT;
  data: any;
}

function isLocalStorageSupport() {
  const test = "test";

  try {
    localStorage.setItem(test, test);
    localStorage.removeItem(test);
    return true;
  } catch (e) {
    return false;
  }
}

export default class Service extends EventEmitter {
  socket: WebSocket;

  token?: Token;

  constructor(host: string) {
    super();

    this.socket = new WebSocket(host);
  }

  async connect(token?: Token) {
    const isSupport = isLocalStorageSupport();

    if (isSupport) {
      //
      token = token || localStorage.getItem("token") || undefined;
      if (!token) {
        throw new Error(`service required token, please connect first`);
      }

      localStorage.setItem("token", token);
    }

    this.token = `Bearer ${token}`;

    await new Promise((resolve) => (this.socket.onopen = resolve));

    this.socket.onmessage = (event) => this.onMessage(event);

    return login(this);
  }

  send(data: Frame) {
    //
    if (!this.token) {
      throw new Error(`service required token, please connect first`);
    }

    const token = this.token;

    this.socket.send(
      btoa(
        JSON.stringify({
          token,
          ...data,
        })
      )
    );
  }

  onMessage(event: MessageEvent) {
    if (!this.token) {
      throw new Error(`service required token, please connect first`);
    }

    const message = JSON.parse(atob(event.data)) as Frame;

    const handler = MUX[message.cmd];

    if (handler) handler(this, message.data);
  }
}
