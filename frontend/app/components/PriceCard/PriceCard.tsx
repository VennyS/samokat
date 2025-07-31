import { CardProps } from "./types";
import styles from "./PriceCard.module.css";
import Image from "next/image";
import { PlusIcon } from "@/app/assets/icons/plus";

export const PriceCard = ({
  href,
  img,
  percentDiscount,
  title,
  weight,
  oldPrice,
  newPrice,
  extraText,
}: CardProps) => {
  return (
    <a href={href}>
      <div className={styles.cardWrapper}>
        <div className={styles.imageWrapper}>
          <Image src={img} alt={title} width={140} height={140} />
          <div className={styles.percentDiscountWrapper}>
            -{percentDiscount}%
          </div>
        </div>
        <p className={styles.title}>{title}</p>
        <div className={styles.weightWrapper}>
          <span>{weight}</span>
          {extraText && (
            <>
              <span>·</span>
              <span className={styles.extraText}>{extraText}</span>
            </>
          )}
        </div>
        <div className={styles.productCardActions}>
          <button>
            <span className={styles.oldPrice}>{oldPrice}</span>
            <span className={styles.newPrice}>{newPrice} ₽</span>
            <PlusIcon />
          </button>
        </div>
      </div>
    </a>
  );
};
