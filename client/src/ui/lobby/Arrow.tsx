import React, { PropsWithChildren, HTMLAttributes } from "react";
import styles from "./Lobby.module.scss";
import clsx from "clsx";
import RES from "../../assets";

type Props = PropsWithChildren<HTMLAttributes<HTMLDivElement>>;

export default function Arrow({ className, style, onClick }: Props) {
  return (
    <div
      className={clsx(styles.control, className)}
      style={style}
      onClick={onClick}
    >
      <img src={RES.getBase64("ARROW")} alt={"ARROW"} />
    </div>
  );
}
