// @ts-nocheck
import {Router} from "express";

const router = Router();

router.get('/v1', (req, res) => {
    res.status(200).json({message: "Work Properly"})
})

export default router;