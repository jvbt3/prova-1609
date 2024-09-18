import 'package:flutter/material.dart';
import 'package:provaesoft8s/service/service.dart';

class UpdateView extends StatefulWidget {
  final int userId;
  final String name;
  final String dataChegada;
  final String dataSaida;
  final int valor;

  UpdateView({
    super.key,
    required this.userId,
    required this.name,
    required this.dataChegada,
    required this.dataSaida,
    required this.valor,
  });

  @override
  State<UpdateView> createState() => _UpdateViewState();
}

class _UpdateViewState extends State<UpdateView> {
  final ViagemService _viagemService = ViagemService();
  late TextEditingController nameController;
  late TextEditingController dataSaidaController;
  late TextEditingController dataChegadaController;
  late TextEditingController valorController;
  late TextEditingController idController;

  @override
  void initState() {
    super.initState();
    nameController = TextEditingController(text: widget.name);
    dataSaidaController = TextEditingController(text: widget.dataSaida);
    dataChegadaController = TextEditingController(text: widget.dataChegada);
    nameController = TextEditingController(text: widget.name);
    valorController = TextEditingController(text: widget.valor.toString());
    idController = TextEditingController(text: widget.userId.toString());
  }

  @override
  void dispose() {
    nameController.dispose();
    idController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Atualizar usuário'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: idController,
                decoration: const InputDecoration(labelText: 'Id'),
                readOnly: true,
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: nameController,
                decoration: const InputDecoration(labelText: 'Nome'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: dataSaidaController,
                decoration: const InputDecoration(labelText: 'Data saida'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: dataChegadaController,
                decoration: const InputDecoration(labelText: 'Data chegada'),
              ),
            ),
            SizedBox(
              width: 150,
              child: TextFormField(
                controller: valorController,
                decoration: const InputDecoration(labelText: 'Valor'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(8.0),
              child: ElevatedButton(
                onPressed: () async {
                  await _viagemService.updateData(
                    id: widget.userId,
                    name: nameController.text,
                    dataChegada: dataChegadaController.text,
                    dataSaida: dataSaidaController.text,
                    valor: double.parse(valorController.text),
                  );
                  Navigator.pop(context);
                },
                child: const Text('Enviar'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
