import React, { useCallback } from "react";
import styles from "./Loading.module.scss";
import { useNavigate } from "react-router-dom";
import RES from "../../assets";

export default function Loading() {
  const navTo = useNavigate();

  const onClick = useCallback(() => {
    navTo(`${process.env.PUBLIC_URL}/lobby`, { replace: true });
  }, [navTo]);

  return (
    <div className={styles.layout} onClick={onClick}>
      <div>
        <img
          className={styles.background}
          src={RES.getBase64("BG")}
          alt={"BG"}
        />
        <img className={styles.logo} src={RES.getBase64("LOGO")} alt={"LOGO"} />

        <div className={styles.click}>
          <h4>press anywhere to start</h4>
        </div>
      </div>
    </div>
  );
}
