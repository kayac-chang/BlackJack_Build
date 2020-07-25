import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import App from './App';
import Game from './game';
import { i18n, gsap } from './plugins';
import service from './service';
import './index.scss';
import store from './store';
import { getURLParam, isLocalStorageSupport } from './utils';
import RES from './assets';
import { PRELOAD, ASSETS } from './assets/pkg';

function getToken() {
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

function getLobby() {
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

function clearURLParam() {
  return window.history.pushState(undefined, '', window.location.origin + window.location.pathname);
}

async function main() {
  await Promise.all([
    //
    i18n.init(),
    gsap.init(),
    service.init(getToken()),
    RES.load(PRELOAD),
  ]);

  getLobby();
  clearURLParam();

  if (process.env.NODE_ENV === 'development') {
    await RES.load(ASSETS);
  }

  const Root = (
    <React.StrictMode>
      <Provider store={store}>
        <App game={Game} />
      </Provider>
    </React.StrictMode>
  );
  ReactDOM.render(Root, document.getElementById('root'));
}

main();
