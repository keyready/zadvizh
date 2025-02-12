import { LoginButton, TelegramAuthData } from '@telegram-auth/react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useCallback, useEffect, useMemo, useState } from 'react';
import { Spinner } from '@heroui/react';

export const MainPage = () => {
    const navigate = useNavigate();
    const [params, setParams] = useSearchParams();

    const [isUserRegistered, setIsUserRegistered] = useState<string>('unknown');
    const [isLoading, setIsLoading] = useState<boolean>(false);

    const handleSuccessAuth = useCallback(async (data: TelegramAuthData) => {
        setIsLoading(true);

        try {
            const result = await fetch(`http://localhost:5000/api/v1/get_access?tgId=${data.id}`);

            if (!result.ok) {
                throw new Error(`HTTP error! Status: ${result.status}`);
            }

            const responseData = await result.json();
            const accessToken = responseData.accessToken;
            localStorage.setItem('t', accessToken);
            navigate('/hierarchy');
        } catch (e) {
            setIsUserRegistered('unregistered');
            setParams({ error: 'need_invitation' });
        } finally {
            setIsLoading(false);
        }
    }, []);

    useEffect(() => {
        const checkReferralValidity = async () => {
            try {
                const result = await fetch(
                    `http://localhost:5000/api/v1/check_ref?ref=${params.get('ref')}`,
                );

                if (!result.ok) {
                    throw new Error(`HTTP error! Status: ${result.status}`);
                }
            } catch (e) {
                setIsUserRegistered('unregistered');
                setParams({ error: 'invalid_referral' });
            } finally {
                setIsLoading(false);
            }
        };

        if (params.get('ref')) {
            setIsLoading(true);
            checkReferralValidity();
        }
    }, [params]);

    const renderErrorMessage = useMemo(() => {
        switch (params.get('error')) {
            case 'need_invitation': {
                return (
                    <h3 className="mb-10 text-center text-xl text-black dark:text-white">
                        Но сначала Вам нужно запросить <br /> приглашение у действующего участника
                        Движа
                    </h3>
                );
            }
            case 'invalid_referral': {
                return (
                    <h3 className="mb-10 text-center text-xl text-black dark:text-white">
                        Но ссылка-приглашение недействительна. <br /> Запросите новую у участника
                        Движа
                    </h3>
                );
            }
            default: {
                return (
                    <h3 className="mb-10 text-center text-xl text-black dark:text-white">
                        Что-то сломалось... Перезагрузите страницу или повторите позже
                    </h3>
                );
            }
        }
    }, [params]);

    if (isUserRegistered === 'unregistered' || params.get('error')) {
        return (
            <section className="bg-main-gradient flex h-screen w-full items-center justify-center">
                <div className="flex min-h-64 min-w-[45%] flex-col items-center gap-3 rounded-xl bg-primary bg-opacity-40 p-4">
                    <h2 className="mt-10 text-center text-2xl font-bold text-black dark:text-white">
                        Увы!
                    </h2>
                    <hr className="w-4/5 opacity-40" />
                    {renderErrorMessage}
                </div>
            </section>
        );
    }

    if (params.get('ref')) {
        return (
            <section className="bg-main-gradient relative flex h-screen w-full items-center justify-center">
                <div className="flex min-h-64 min-w-[33%] flex-col items-center rounded-xl bg-primary bg-opacity-40 p-4">
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
                {isLoading && (
                    <div className="absolute bottom-0 left-0 right-0 top-0 flex items-center justify-center bg-primary bg-opacity-50 backdrop-blur">
                        <div className="flex h-32 w-64 items-center justify-center rounded-md bg-gray-400 bg-opacity-50">
                            <Spinner size="lg" />
                        </div>
                    </div>
                )}
            </section>
        );
    }

    return (
        <section className="bg-main-gradient relative flex h-screen w-full items-center justify-center">
            <div className="flex min-h-64 min-w-[33%] flex-col items-center justify-center gap-4 rounded-xl bg-primary bg-opacity-40 p-4">
                <h1 className="text-center text-2xl font-bold text-black dark:text-white">
                    Для продолжения работы <br /> авторизуйтесь
                </h1>
                <LoginButton
                    cornerRadius={10}
                    onAuthCallback={handleSuccessAuth}
                    botUsername={import.meta.env.VITE_BOT_USERNAME}
                />
            </div>
            {isLoading && (
                <div className="absolute bottom-0 left-0 right-0 top-0 flex items-center justify-center bg-primary bg-opacity-50 backdrop-blur">
                    <div className="flex h-32 w-64 items-center justify-center rounded-md bg-gray-400 bg-opacity-50">
                        <Spinner size="lg" />
                    </div>
                </div>
            )}
        </section>
    );
};
