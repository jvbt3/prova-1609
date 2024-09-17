import 'dart:convert';
import 'package:http/http.dart' as http;

class UserService {
  Future<http.Response> postData(
      {required String name,
      required int id,
      required String dataChegada,
      required String dataSaida,
      required double valor}) async {
    try {
      final response = await http.post(
        Uri.parse('http://localhost:6666/api/viagem'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, dynamic>{
          'name': name,
          'id': id,
          'data_chegada': dataChegada,
          'data_saida': dataSaida,
          'valor': valor
        }),
      );

      return response;
    } catch (e) {
      throw Exception('Error: $e');
    }
  }
}
