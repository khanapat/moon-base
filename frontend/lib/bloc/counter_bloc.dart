// import 'package:bloc/bloc.dart';

// enum CounterEvent { increment, decrement }

// class CounterState {
//   final int value;
//   CounterState({required this.value});

//   @override
//   String toString() {
//     // TODO: implement toString
//     return 'Value is ${value}';
//   }
// }

// class CounterBloc extends Bloc<CounterEvent, CounterState> {
//   CounterBloc() : super(CounterState(value: 0));

//   @override
//   Stream<CounterState> mapEventToState(CounterEvent event) async* {
//     // TODO: implement mapEventToState
//     switch (event) {
//       case CounterEvent.increment:
//         yield CounterState(value: state.value + 1);
//         break;
//       case CounterEvent.decrement:
//         yield CounterState(value: state.value - 1);
//         break;
//     }
//   }
// }