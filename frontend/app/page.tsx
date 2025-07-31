import styles from "./page.module.css";
import { PriceCard } from "./components/PriceCard/PriceCard";

export default function Home() {
  return (
    <main className={styles.main}>
      <PriceCard
        href="/product"
        img="/beacon.jpg"
        percentDiscount="26"
        title="Сырокопчёный бекон Дымов, нарезка"
        weight="150 г"
        oldPrice="229"
        newPrice="169"
      />
      <PriceCard
        href="/product"
        img="/monterra.jpg"
        percentDiscount="34"
        title="Мороженое Monterra, с кленовым сиропом и грецким орехом, в ведёрке"
        weight="298 г"
        extraText="Снизили цену"
        oldPrice="749"
        newPrice="494"
      />
    </main>
  );
}
