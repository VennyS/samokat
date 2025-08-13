import { Props } from "./types";
import cn from "classnames";
import styles from "./Button.module.css";

export const Button = ({ title, img, onClick, className }: Props) => {
  return (
    <button onClick={onClick} className={cn(styles.button, className)}>
      {img && <img src={img} alt="" className={styles.icon} />}
      {title && <span className={styles.span}>{title}</span>}
    </button>
  );
};
