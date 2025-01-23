"use client";
import { getLinks, Shortlink, submitLink } from "@exxo/requests";
import { useMutation, useQuery } from "@tanstack/react-query";
import styles from "./main.module.css";

export const Main = () => {
  const {
    data,
    error,
    refetch: refetchLinks,
  } = useQuery({
    queryKey: ["shortlinks"],
    queryFn: getLinks,
  });

  const { mutateAsync: submitLongUrl } = useMutation({
    mutationKey: ["shortlink"],
    mutationFn: submitLink,
  });

  const onSubmit = async (e) => {
    e.preventDefault();
    const input = e.target.elements?.[0];
    if (input.value) {
      await submitLongUrl(input.value);
      await refetchLinks();
      input.value = "";
    }
  };
  const onCopy = async (url: string) => {
    await navigator.clipboard.writeText(url);
  };
  return (
    <div className={styles.wrapper}>
      <div className={styles.title}>The Linker</div>
      <form className={styles.submitbox} onSubmit={onSubmit}>
        <input className={styles.input} />
        <input className={styles.submit} type="submit" value="Short" />
      </form>
      <div className={styles.links}>
        {data &&
          (data as unknown as Shortlink[])
            .sort(
              (x, y) => new Date(y.date).getTime() - new Date(x.date).getTime(),
            )
            .map((item) => (
              <div className={styles.linkwrapper} key={item.short_url}>
                {void console.log("item", item)}
                <a
                  className={styles.link}
                  target="_blank"
                  href={item.short_url}
                >
                  {item.short_url}
                </a>
                <div onClick={() => onCopy(item.short_url)}>
                  <img
                    className={styles.copy}
                    src="/clipboard.png"
                    alt="copy"
                  />
                </div>
              </div>
            ))}
      </div>
    </div>
  );
};
