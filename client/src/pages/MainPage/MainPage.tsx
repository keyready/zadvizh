import { LoginButton } from '@telegram-auth/react';
import { useNavigate, useSearchParams } from 'react-router-dom';

export const MainPage = () => {
    const navigate = useNavigate();
    const [params] = useSearchParams();

    if (!params.get('ref')) {
        return (
            <section className="bg-main-gradient flex h-screen w-full items-center justify-center">
                <div className="min-w-1/3 flex min-h-64 flex-col items-center justify-center rounded-xl bg-primary p-4">
                    <h1 className="text-center text-3xl font-bold text-black dark:text-white">
                        Без приглашения сюда не попасть :(
                    </h1>
                </div>
            </section>
        );
    }

    return (
        <section className="bg-main-gradient flex h-screen w-full items-center justify-center">
            <div className="min-w-1/3 flex min-h-64 flex-col items-center rounded-xl bg-primary p-4">
                <h1 className="text-center text-3xl font-bold text-black dark:text-white">
                    Добро пожаловать в Движ
                </h1>
                <h2 className="mt-10 text-center text-2xl text-black dark:text-white">
                    Вы попали сюда не случайно!
                </h2>
                <h3 className="mb-10 text-center text-xl text-black dark:text-white">
                    Для получения доступа к чату, авторизуйтесь через Telegram
                </h3>
                <LoginButton
                    cornerRadius={10}
                    onAuthCallback={(data) => {
                        console.log(data);
                        navigate(`/auth/continue?ref=${params.get('ref') || ''}&un=${data.id}`);
                    }}
                    botUsername={import.meta.env.VITE_BOT_USERNAME}
                />
            </div>
        </section>
    );
};
