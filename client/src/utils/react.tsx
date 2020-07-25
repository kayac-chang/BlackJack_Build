import React, { ReactNode, createContext, useReducer, useContext } from 'react';

type Props = { children: ReactNode };
type Dispatch<Action> = (action: Action) => void;
type Reducer<State, Action> = (state: State, action: Action) => State;

export function ContextProvider<State, Action>(initialState: State, reducer: Reducer<State, Action>) {
  //
  const StateContext = createContext<State | undefined>(initialState);
  const DispatchContext = createContext<Dispatch<Action> | undefined>(undefined);

  function Provider({ children }: Props) {
    const [state, dispatch] = useReducer(reducer, initialState);

    return (
      <StateContext.Provider value={state}>
        <DispatchContext.Provider value={dispatch}>{children}</DispatchContext.Provider>
      </StateContext.Provider>
    );
  }

  function useState() {
    const state = useContext(StateContext);
    const dispatch = useContext(DispatchContext);

    if (!state || !dispatch) {
      throw new Error('must be used within a Provider');
    }

    return { state, dispatch };
  }

  return { Provider, useState };
}
