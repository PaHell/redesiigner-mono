/// <reference types="@sveltejs/kit" />

type TButton<TIn, TOut> = {
    icon?: string | undefined;
    text?: string | undefined;
    color?: string | undefined;
    OnClick: (TIn) => TOut;
};

type TBanner = {
    icon: string;
    text: string;
    color?: string;
    buttons: Button<void, void>[];
    remove?: () => void;
};

type TModal = {
    icon: string;
    title: string;
    text: string;
    color: string;
    buttons: Button<(close: () => void) => void, void>[];
    remove?: () => void;
};

type TNotification = {
    icon?: string;
    img?: string;
    title: string;
    text: string;
    color: string;
    variant?: "default" | "button" | "reply";
    buttons: Button<void, void>[];
    remove?: () => void;
};