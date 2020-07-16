import { createStore, applyMiddleware } from 'redux';
import thunkMiddleware from 'redux-thunk';
import { composeWithDevTools } from 'redux-devtools-extension';
import reducer from './reducers';

const middlewares = [thunkMiddleware];

const store = createStore(
  //
  reducer,
  composeWithDevTools(applyMiddleware(...middlewares))
);

type SelectFunc<T> = (state: AppState) => T;
type OnChangeFunc<T> = (state: T) => void;

function observe<T>(select: SelectFunc<T>, onChange: OnChangeFunc<T>) {
  //
  let previous: T;

  function handleChange() {
    const current = select(store.getState());

    if (previous !== current) {
      previous = current;

      onChange(current);
    }
  }

  return store.subscribe(handleChange);
}

export type AppState = ReturnType<typeof reducer>;
export { observe };
export default store;
