import {Input} from "@heroui/input";
import {Button} from "@heroui/button";
import {Progress, Radio, RadioGroup} from "@heroui/react";
import {FormEvent, useCallback, useEffect, useMemo, useState} from "react";
import {useParams} from "react-router";

type FieldType = "dev" | "sec" | "devops" | "science" | "org"

interface User {
    firstname: string;
    lastname: string;
    department: string;
    field: FieldType;
    teamName?: string;
    teamRole?: string;
    position?: string;
    ref?: string;
    scidir: string;
}

export const RegisterPage = () => {
    const {ref} = useParams<{ ref: string }>();

    const [step, setStep] = useState<number>(1);
    const [newUserForm, setNewUserForm] = useState<Partial<User>>({});
    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [commPassword, setCommPassword] = useState<string>("");

    useEffect(() => {
        setNewUserForm(ps => ({
            ...ps,
            ref
        }));
    }, [ref]);

    const handleStepChange = useCallback(async (ev: FormEvent<HTMLFormElement>) => {
        ev.preventDefault();

        if (step === 3 && newUserForm.field === "devops") {
            if (commPassword === import.meta.env.VITE_DEVOPS_PASS) {
                setStep(ps => ps + 1);
            } else return;
        }

        if (step === 3 && newUserForm.field === "org") {
            if (commPassword === import.meta.env.VITE_COMMANDERS_PASS) {
                setStep(ps => ps + 1);
            } else return;
        }

        if (step < 4) setStep(ps => ps + 1);
        if (step === 4) {
            console.log('print');
            setIsLoading(true);
            try {
                const result = await fetch("/api/v1/auth", {
                    body: JSON.stringify(newUserForm),
                    method: "post",
                    headers: {
                        "Content-Type": "application/json"
                    }
                });

                if (result.ok) {
                    window.location.href = import.meta.env.VITE_ZADVIZH_LINK;
                }
            } catch (e) {
                alert(e)
            } finally {
                setIsLoading(false);
            }
        }
    }, [commPassword, newUserForm, step]);

    const showButtonBeDisabled = useMemo(() => {
        switch (step) {
            case 1: {
                return !newUserForm.firstname || !newUserForm.lastname || !newUserForm.department;
            }
            case 2: {
                return !newUserForm.field;
            }
            default: {
                if (newUserForm.field === "dev" || newUserForm.field === "sec") {
                    return !newUserForm.position;
                }
                if (newUserForm.field === "devops") {
                    return newUserForm.scidir !== import.meta.env.VITE_DEVOPS_PASS;
                }
                if (newUserForm.field === "org") {
                    return commPassword !== import.meta.env.VITE_COMMANDERS_PASS;
                }
                if (newUserForm.field === "science") {
                    return !newUserForm.scidir;
                }
            }
        }
    }, [commPassword, newUserForm.department, newUserForm.field, newUserForm.firstname, newUserForm.lastname, newUserForm.position, newUserForm.scidir, step]);

    const renderQuestions = useMemo(() => {
        switch (step) {
            case 1: {
                return (
                    <>
                        <Input
                            onValueChange={val => setNewUserForm(ps => ({
                                ...ps,
                                firstname: val
                            }))}
                            value={newUserForm.firstname}
                            label="Ваше имя"
                        />
                        <Input
                            onValueChange={val => setNewUserForm(ps => ({
                                ...ps,
                                lastname: val
                            }))}
                            value={newUserForm.lastname}
                            label="Ваша фамилия"
                        />
                        <Input
                            onValueChange={val => setNewUserForm(ps => ({
                                ...ps,
                                department: val
                            }))}
                            value={newUserForm.department}
                            label="Отдел"
                        />
                    </>
                );
            }
            case 2: {
                return (
                    <>
                        <RadioGroup
                            value={newUserForm.field}
                            onValueChange={(val) => setNewUserForm(ps => ({
                                ...ps,
                                field: val as FieldType
                            }))}
                            color="success"
                            label="Выберите основное направление деятельности"
                        >
                            <Radio value="dev">Разработка</Radio>
                            <Radio value="sec">Безопасность</Radio>
                            <Radio value="devops">DevOps</Radio>
                            <Radio value="science">Научная деятельность</Radio>
                            <Radio value="org">Организация рабочего процесса</Radio>
                        </RadioGroup>
                    </>
                );
            }
            case 3: {
                return (
                    <>
                        {(newUserForm.field === "dev" || newUserForm.field === "sec") && (
                            <div>
                                <h3 className="">Вы состоите в какой-нибудь команде?</h3>
                                <Input
                                    label="Название команды (или оставьте пустым)"
                                    onValueChange={val => setNewUserForm(ps => ({
                                        ...ps,
                                        teamName: val
                                    }))}
                                    value={newUserForm.teamName}
                                />

                                <h3 className="mt-3">Ваша роль в команде</h3>
                                <RadioGroup
                                    isDisabled={!newUserForm.teamName}
                                    onValueChange={val => setNewUserForm(ps => ({
                                        ...ps,
                                        teamRole: val
                                    }))}
                                    value={newUserForm.teamRole}
                                    color="success"
                                >
                                    <Radio value="cap">Капитан</Radio>
                                    <Radio value="part">Участник</Radio>
                                </RadioGroup>

                                <h3 className="mt-3">Ваше направление</h3>
                                <RadioGroup
                                    onValueChange={val => setNewUserForm(ps => ({
                                        ...ps,
                                        position: val
                                    }))}
                                    value={newUserForm.position}
                                    color="success"
                                >
                                    {newUserForm.field === "dev" ? (
                                        <>
                                            <Radio value="front">Frontend</Radio>
                                            <Radio value="back">Backend</Radio>
                                            <Radio value="ml">Machine Learning</Radio>
                                            <Radio value="design">UI/UX</Radio>
                                        </>
                                    ) : (
                                        <>
                                            <Radio value="web">Websec</Radio>
                                            <Radio value="pwn">PWN</Radio>
                                            <Radio value="froensic">Форензика</Radio>
                                            <Radio value="admin">Администрирование</Radio>
                                            <Radio value="crypto">Криптографию</Radio>
                                            <Radio value="stegano">Стеганография</Radio>
                                            <Radio value="osint">OSINT</Radio>
                                            <Radio value="joy">Игрушки</Radio>
                                        </>
                                    )}
                                </RadioGroup>
                            </div>
                        )}

                        {newUserForm.field === "devops" && (
                            <Input
                                onValueChange={val => setNewUserForm(ps => ({
                                    ...ps,
                                    scidir: val
                                }))}
                                value={newUserForm.scidir}
                                label="Кто у нас старший за сети (Фамилия И.О.)?"
                            />
                        )}

                        {newUserForm.field === "science" && (
                            <Input
                                onValueChange={val => setNewUserForm(ps => ({
                                    ...ps,
                                    scidir: val
                                }))}
                                value={newUserForm.scidir}
                                label="Ваш научрук"
                            />
                        )}

                        {newUserForm.field === "org" && (
                            <Input
                                type="password"
                                onValueChange={setCommPassword}
                                value={commPassword}
                                label="Специальный пароль"
                            />
                        )}
                    </>
                );
            }
            default: {
                return (
                    <div className="w-full mt-10 mb-3">
                        <h1 className="text-center text-xl font-bold">Регистрация завершена!</h1>
                        <h3 className="text-center">Нажмите "Завершить", чтобы получить приглашение в группу</h3>
                    </div>
                );
            }
        }
    }, [commPassword, newUserForm.department, newUserForm.field, newUserForm.firstname, newUserForm.lastname, newUserForm.position, newUserForm.scidir, newUserForm.teamName, newUserForm.teamRole, step]);

    return (
        <section className="flex items-center justify-center w-full h-screen bg-main-gradient">
            <div className="flex flex-col items-center p-4 w-1/3 min-h-1/3 bg-primary rounded-xl">

                {step <= 3 && (
                    <>
                        <Progress color="success" className="max-w-md" value={step * (100 / 3)}/>
                        <h1 className="text-center mt-5  font-bold text-black dark:text-white text-3xl">
                            Продолжение регистрации
                        </h1>
                    </>
                )}

                <form onSubmit={handleStepChange} className="flex flex-col gap-4">
                    {step <= 3 && (
                        <h1>Заполните поля формы, чтобы завершить регистрацию</h1>
                    )}

                    {renderQuestions}

                    <div className="flex justify-end gap-2">
                        {step > 1 && step !== 4 && (
                            <Button
                                type="button"
                                onPress={() => setStep(ps => ps - 1)}
                                className="w-1/5"
                            >
                                Назад
                            </Button>
                        )}
                        <Button
                            isDisabled={showButtonBeDisabled}
                            isLoading={isLoading}
                            type="submit"
                            className="w-1/3"
                            color="success"
                        >
                            {step < 3 ? "Продолжить" : "Завершить!"}
                        </Button>
                    </div>
                </form>
            </div>
        </section>
    );
};
