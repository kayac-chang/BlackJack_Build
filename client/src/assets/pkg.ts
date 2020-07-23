import CHIP_FLAT_RED from './chip/flat/red.png';
import CHIP_FLAT_GREEN from './chip/flat/green.png';
import CHIP_FLAT_BLUE from './chip/flat/blue.png';
import CHIP_FLAT_BLACK from './chip/flat/black.png';
import CHIP_FLAT_PURPLE from './chip/flat/purple.png';
import CHIP_FLAT_YELLOW from './chip/flat/yellow.png';

import CHIP_NORMAL_RED from './chip/normal/red.png';
import CHIP_NORMAL_GREEN from './chip/normal/green.png';
import CHIP_NORMAL_BLUE from './chip/normal/blue.png';
import CHIP_NORMAL_BLACK from './chip/normal/black.png';
import CHIP_NORMAL_PURPLE from './chip/normal/purple.png';
import CHIP_NORMAL_YELLOW from './chip/normal/yellow.png';

import ICON_BJ from './icon/BJ.png';
import ICON_BUST from './icon/BUST.png';
import ICON_LOSE from './icon/LOSE.png';
import ICON_WIN from './icon/WIN.png';
import ICON_DEAL from './icon/DEAL.png';

import ARROW from './lobby/arrow.png';
import BG from './lobby/background.jpg';
import DETAIL from './lobby/detail.png';
import LOGO from './lobby/logo.png';
import TABLE from './lobby/table.png';
import ROOM_NUM from './lobby/room_number.png';

import SELECT_SEAT_NORMAL from './seat/MultiSeat.png';
import SELECT_SEAT_ENABLE from './seat/MultiSeat_On.png';
import SEAT_NORMAL from './seat/Seat.png';
import SEAT_ENABLE from './seat/Seat_On.png';
import FIELD from './seat/Field.png';
import JOIN from './seat/Join.png';

import TABLES from './table';
import POKER from './poker';
import { BGM, SFX } from './sound';

const PRELOAD = Object.freeze({
  BG,
  LOGO,
});

const ASSETS = Object.freeze({
  CHIP_FLAT_RED,
  CHIP_FLAT_GREEN,
  CHIP_FLAT_BLUE,
  CHIP_FLAT_BLACK,
  CHIP_FLAT_PURPLE,
  CHIP_FLAT_YELLOW,
  //
  CHIP_NORMAL_RED,
  CHIP_NORMAL_GREEN,
  CHIP_NORMAL_BLUE,
  CHIP_NORMAL_BLACK,
  CHIP_NORMAL_PURPLE,
  CHIP_NORMAL_YELLOW,
  //
  ICON_BJ,
  ICON_BUST,
  ICON_LOSE,
  ICON_WIN,
  ICON_DEAL,
  //
  ARROW,
  DETAIL,
  TABLE,
  ROOM_NUM,
  //
  SELECT_SEAT_NORMAL,
  SELECT_SEAT_ENABLE,
  SEAT_NORMAL,
  SEAT_ENABLE,
  FIELD,
  JOIN,
  //
  ...TABLES,
  ...POKER,
  //
  ...BGM,
  ...SFX,
});

export type PKG = keyof typeof PRELOAD | keyof typeof ASSETS;

export { PRELOAD, ASSETS };
