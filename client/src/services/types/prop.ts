export interface RoomProp {
  id: number;
  max_bet: number;
  min_bet: number;
  history: (number | string)[];
  occupied: number;
}

export interface SeatProp {
  no: number;
  player: string;
  total_bet: number;
  pay?: number;
  piles?: PileProp[];
}

export interface PileProp {
  bet: number;
  cards: string[];
  pay: number;
  result: number;
}

export type GameStateProp = [number, number?, number?];

export interface GameProp {
  id: number;
  round: string;
  state: GameStateProp;
  seats: SeatProp[];
  shoe_num: number;
  max_bet: number;
  min_bet: number;
}

export interface CountDownProp {
  expire: number;
}

export interface DealProp {
  card: string;
  cards: string[];
  no: number;
  pile: number;
  points: number;
  shoe_num: number;
}

export interface LoginProp {
  user_name: string;
}

export interface UpdateProp {
  name: string;
  balance: number;
}

export interface OptionsProp {
  dbl: boolean;
  gvp: boolean;
  hit: boolean;
  ins: boolean;
  pay: boolean;
  spt: boolean;
  sty: boolean;
}

export interface TurnProp {
  no: number;
  pile: number;
  expire: number;
  options: OptionsProp;
}

export interface JoinSeatProp {
  id: number;
  no: number;
}

export interface ActionProp {
  action: string;
  id: number;
  no: number;
  pile: number;
}
