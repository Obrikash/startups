import { EyeIcon } from "lucide-react";
import React from "react";

const View = ({ views }: { views: number }) => {
  return (
    <div className="view-container">
      <p className="view-text">
        <span className="font-black">
          {views} {<EyeIcon className="size-6 text-primary-default" />}
        </span>
      </p>
    </div>
  );
};

export default View;
