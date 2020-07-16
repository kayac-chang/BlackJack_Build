import React, { PropsWithChildren } from "react";
import styles from "./GameRules.module.scss";
import { useTranslation } from "react-i18next";

type SectionProps = {
  title: string;
  descriptions: string[];
};
function Section({ title, descriptions }: SectionProps) {
  return (
    <section>
      <h4>{title}</h4>

      <ul>
        {descriptions.map((description, index) => (
          <li key={`${title}${index}`}>{description}</li>
        ))}
      </ul>
    </section>
  );
}

type TitleProps = PropsWithChildren<{}>;
function Title({ children }: TitleProps) {
  return (
    <section>
      <div>
        <h3>{children}</h3>
      </div>
    </section>
  );
}

export default function GameRules() {
  const [t] = useTranslation();

  const sections = t("gamerules", { returnObjects: true }) as SectionProps[];

  return (
    <div className={styles.gameRules}>
      <Title>game rules</Title>

      {sections.map(({ title, descriptions }) => (
        <Section key={title} title={title} descriptions={descriptions} />
      ))}
    </div>
  );
}
