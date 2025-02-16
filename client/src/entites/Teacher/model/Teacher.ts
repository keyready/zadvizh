import { Comment } from '../../Comment';

export interface Teacher {
    id: string;
    firstname: string;
    lastname: string;
    middlename: string;
    likes: { value: number; authors: string[] };
    dislikes: { value: number; authors: string[] };
    comments: Comment[];
}
