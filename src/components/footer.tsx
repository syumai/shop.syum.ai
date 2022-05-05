import Link from "next/link";
import { VFC } from "react";

import { footerStyle } from "../style/footer.css";

export const Footer: VFC = () => (
  <div className={footerStyle.wrapper}>
    <div className={footerStyle.innerWrapper}>
      <Link href="/about">
        <a className={footerStyle.link}>このサイトについて</a>
      </Link>
      <p>
        created by{" "}
        <a
          href="https://twitter.com/__syumai"
          className={footerStyle.link}
        >
          syumai
        </a>
      </p>
      <p>This site use cookie for Google Analytics</p>
      <p>
        Source code is&nbsp;
        <a
          href="https://github.com/syumai/shop.syum.ai"
          className={footerStyle.link}
        >
          here
        </a>
      </p>
    </div>
  </div>
);
