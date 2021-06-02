export function isLocalStorageSupport() {
  const test = 'test';

  try {
    localStorage.setItem(test, test);
    localStorage.removeItem(test);
    return true;
  } catch (e) {
    return false;
  }
}

export function getURLParam(key: string) {
  const url = new URL(window.location.href);

  return url.searchParams.get(key) || undefined;
}

export function getToken() {
  if (!isLocalStorageSupport()) {
    return;
  }

  const token = getURLParam('token') || localStorage.getItem('token') || undefined;

  if (!token) {
    throw new Error(`can not access without token, please contact the service.`);
  }

  localStorage.setItem('token', token);

  return token;
}

export function getLobby() {
  if (!isLocalStorageSupport()) {
    return;
  }

  const lobby = getURLParam('lobby') || localStorage.getItem('lobby') || undefined;

  if (!lobby) {
    return;
  }

  localStorage.setItem('lobby', lobby);

  return lobby;
}

export function getDebug() {
  if (!isLocalStorageSupport()) {
    return;
  }

  const debug = getURLParam('debug') || localStorage.getItem('debug') || undefined;

  if (!debug) {
    return;
  }

  localStorage.setItem('debug', debug);

  return debug === 'true';
}

export function clearURLParam() {
  return window.history.pushState(undefined, '', window.location.origin + window.location.pathname);
}
