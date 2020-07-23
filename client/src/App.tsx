import React, { useEffect, useState } from 'react';
import { Canvas } from './ui/components';
import { BrowserRouter as Router, Routes, Route, useParams, Navigate } from 'react-router-dom';
import services from './service';
import Loading from './ui/loading';
import { ModalProvider } from './ui/modal';
import { SoundProvider } from './sound';
import Frame from './ui/Frame';
import Lobby from './ui/lobby';

type Props = {
  game: (canvas: HTMLCanvasElement) => void;
};

function useJoin() {
  const params = useParams();
  const [hasJoin, setHasJoin] = useState(false);

  useEffect(() => {
    services.joinRoom(Number(params.id)).then(() => setHasJoin(true));
  }, [params]);

  return hasJoin;
}

function Game({ game }: Props) {
  const join = useJoin();

  if (join) {
    return <Canvas>{game}</Canvas>;
  }

  return <></>;
}

export default function App({ game }: Props) {
  return (
    <ModalProvider>
      <SoundProvider>
        <Router>
          <Frame>
            <Routes basename={process.env.PUBLIC_URL}>
              <Route path="/" element={<Loading />} />
              <Route path="lobby" element={<Lobby />} />
              <Route path="game/:id" element={<Game game={game} />} />
              <Route path="*" element={<Navigate to="/" />} />
            </Routes>
          </Frame>
        </Router>
      </SoundProvider>
    </ModalProvider>
  );
}
