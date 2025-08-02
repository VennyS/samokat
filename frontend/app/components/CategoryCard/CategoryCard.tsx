import { Props } from "./types";
import styles from "./CategoryCard.module.css";

export const CategoryCard = ({ href, img, title }: Props) => {
  return (
    <a href={href}>
      <div className={styles.cardWrapper}>
        <p>{title}</p>
        <img src={img} alt={title} />
      </div>
    </a>
  );
};
