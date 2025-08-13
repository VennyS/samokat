import { Props } from "./types";
import styles from "./SearchField.module.css";
import Image from "next/image";

export const SearchField = ({ placeholder, classname }: Props) => {
  return (
    <div className={styles.searchDiv}>
      <input
        type="text"
        placeholder={placeholder}
        className={styles.searchField}
      />
      <Image src="/search.png" alt="Поиск" height={16} width={16} />
    </div>
  );
};
