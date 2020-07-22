import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import App from './App';
import Game from './game';
import { i18n, gsap } from './plugins';
import service from './service';
import './index.scss';
import store from './store';
import { getURLParam } from './utils';
import RES from './assets';

async function main() {
  await Promise.all([i18n.init(), gsap.init(), service.init(getURLParam('token')), RES.load()]);

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
