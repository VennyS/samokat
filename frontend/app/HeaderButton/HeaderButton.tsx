import { Props } from "./types";
import styles from "./HeaderButton.module.css";
import Image from "next/image";
import cn from "classnames";

export const HeaderButton = ({ title, className, icon }: Props) => {
  return (
    <button className={cn(styles.button, className)}>
      <Image src={icon} alt={title} height={24} width={24}></Image>
      <span>{title}</span>
    </button>
  );
};
