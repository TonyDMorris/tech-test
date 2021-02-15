import {setRender} from '../sharedCode';
import {storeData} from '../API/StoreData';

export const renderReducer = (state, action) => {
  switch (action.type) {
    case 'StoreQuiz':
      console.log("StoreQuiz");
      state.completed = action.update.completed;
      let answers = createResponse(state);
      console.log(answers);

      storeData(answers);

      return state;
    case 'AnswerQuestion':
      state.answers.set(state.currentQuestionId, action.update.answer);
      const currentQuestion = state.questionMap.get(state.currentQuestionId);

      // Move to the next question
      state.currentQuestionId = currentQuestion.nextQuestionId 
                               ? currentQuestion.nextQuestionId : state.answerMap.get(action.update.answer).nextQuestionId;
  
      // call re-render
      const nextQuestion = state.questionMap.get(state.currentQuestionId);
      let nextRender = setRender(nextQuestion);
      action.update.nextPage(nextRender);

      return state;
    default:
      return state;
  }
}

const createResponse = (state) => {
  return {
    quizId: "" + state.quizId,
    questionAnswers: createAnswerObject(state.answers)
  }
}

const createAnswerObject = (answerMap) => {
  let results = [];

  answerMap.forEach((value, key) => {
    let question = {};
    question.id="" + key;
    question.answers = Array.isArray(value) ? value : ["" + value];
    results.push(question);
  });

  return results;
}