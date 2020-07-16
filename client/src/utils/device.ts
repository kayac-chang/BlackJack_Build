import Detect from "mobile-detect";

function Detecter() {
  return new Detect(window.navigator.userAgent);
}

export function isMobile() {
  const detect = Detecter();

  return Boolean(detect.mobile());
}
