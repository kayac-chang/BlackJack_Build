import React, { useCallback, useEffect, useState, CSSProperties } from "react";
import styles from "./Loading.module.scss";
import { useNavigate } from "react-router-dom";
import RES from "../../assets";
import { ASSETS } from "../../assets/pkg";

type Props = {
  max?: number;
  min?: number;
  value: number;
  style?: CSSProperties;
};
function ProgressBar({ max = 100, min = 0, value, style }: Props) {
  const range = max - min;
  const width = (value / range) * 100;

  return (
    <div className={styles.progressbar} style={style}>
      <div className={styles.progressdone} style={{ width: `${width}%` }} />
    </div>
  );
}

export default function Loading() {
  const navTo = useNavigate();

  const [progress, setProgress] = useState(0);

  const onClick = useCallback(() => {
    if (progress < 100) {
      return;
    }

    navTo(`${process.env.PUBLIC_URL}/lobby`, { replace: true });
  }, [progress, navTo]);

  useEffect(() => {
    RES.onProgress(setProgress);

    RES.load(ASSETS).then(() => setProgress(100));
  }, [setProgress]);

  return (
    <div className={styles.layout} onClick={onClick}>
      <div>
        <img
          className={styles.background}
          src={RES.getBase64("BG")}
          alt={"BG"}
        />
        <img className={styles.logo} src={RES.getBase64("LOGO")} alt={"LOGO"} />

        <ProgressBar
          value={progress}
          style={{ opacity: progress < 100 ? 1 : 0 }}
        />

        <div
          className={styles.click}
          style={{ opacity: progress < 100 ? 0 : 1 }}
        >
          <h4>press anywhere to start</h4>
        </div>
      </div>
    </div>
  );
}
